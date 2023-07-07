import { CoreApp, DataQuery, DataSourceApi, hasQueryExportSupport, hasQueryImportSupport } from '@grafana/data';
import { isExpressionReference } from '@grafana/runtime/src/utils/DataSourceWithBackend';
import { getNextRefIdChar } from 'app/core/utils/query';

export async function updateQueries(
  nextDS: DataSourceApi,
  nextDSUidOrVariableExpression: string,
  queries: DataQuery[],
  currentDS?: DataSourceApi
): Promise<DataQuery[]> {
  let nextQueries = queries;
  const datasource = { type: nextDS.type, uid: nextDSUidOrVariableExpression };
  const DEFAULT_QUERY = { ...nextDS?.getDefaultQuery?.(CoreApp.PanelEditor), datasource, refId: 'A' };

  // we are changing data source type
  if (currentDS?.meta.id !== nextDS.meta.id) {
    // If changing to mixed do nothing
    if (nextDS.meta.mixed) {
      return queries;
    }
    // when both data sources support abstract queries
    else if (hasQueryExportSupport(currentDS) && hasQueryImportSupport(nextDS)) {
      const abstractQueries = await currentDS.exportToAbstractQueries(queries);
      nextQueries = await nextDS.importFromAbstractQueries(abstractQueries);
    }
    // when datasource supports query import
    else if (currentDS && nextDS.importQueries) {
      nextQueries = await nextDS.importQueries(queries, currentDS);
    }
    // Otherwise clear queries that do not match the next datasource UID
    else {
      if (currentDS) {
        const reducedQueries: DataQuery[] = [];
        const nextDsQueries = queries.reduce((reduced, currentQuery) => {
          if (currentQuery.datasource && currentQuery.datasource.uid === nextDS.uid) {
            currentQuery.refId = getNextRefIdChar(reduced);
            return reduced.concat([currentQuery]);
          }
          return reduced;
        }, reducedQueries);

        if (nextDsQueries.length > 0) {
          return nextDsQueries;
        }
      }

      return [DEFAULT_QUERY];
    }
  }

  if (nextQueries.length === 0) {
    return [DEFAULT_QUERY];
  }

  // Set data source on all queries except expression queries
  return nextQueries.map((query) => {
    if (!isExpressionReference(query.datasource) && !nextDS.meta.mixed) {
      query.datasource = datasource;
    }
    return query;
  });
}
