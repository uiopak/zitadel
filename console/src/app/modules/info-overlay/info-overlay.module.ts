import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatLegacyButtonModule as MatButtonModule } from '@angular/material/legacy-button';
import { MatIconModule } from '@angular/material/icon';
import { TranslateModule } from '@ngx-translate/core';

import { InfoOverlayComponent } from './info-overlay.component';

@NgModule({
  declarations: [InfoOverlayComponent],
  imports: [CommonModule, MatButtonModule, MatIconModule, TranslateModule],
  exports: [InfoOverlayComponent],
})
export class InfoOverlayModule {}
