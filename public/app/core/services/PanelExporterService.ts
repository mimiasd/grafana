import appEvents from 'app/core/app_events';

import { ExportPanelPayload, PanelExportEvent } from '../../types/events';

export class PanelExporterService {
  init() {
    appEvents.subscribe(PanelExportEvent, (e) => this.exportPNG(e.payload)); // Binds PanelExportEvent to exportPNG method
  }

  exportPNG(e: ExportPanelPayload) {
    console.log('exportPNG', e);
    //todo avoid as     DONE?
    const canvas = e.htmlElement;

    //TODO FIX
    const link = document.createElement('a');
    link.download = 'asdasd.png'; // Do we want custom names or nah?
    canvas.toBlob((blob) => {
      link.href = URL.createObjectURL(blob!);
      link.click();
      URL.revokeObjectURL(link.href);
    }, 'image/png');
  }

  exportCSV(e: ExportPanelPayload) {
    console.log('exportCSV', e);
  }

  /*
  onClick={() => {
              if (hasLogs) {
                reportInteraction('grafana_logs_download_clicked', {
                  app,
                  format: 'csv',
                });
              }
              this.exportCsv(dataFrames[dataFrameIndex], { useExcelHeader: this.state.downloadForExcel });
            }}
            */
}
