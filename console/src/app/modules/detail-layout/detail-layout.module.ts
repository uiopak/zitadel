import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatLegacyButtonModule as MatButtonModule } from '@angular/material/legacy-button';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { BackModule } from 'src/app/directives/back/back.module';

import { DetailLayoutComponent } from './detail-layout.component';

@NgModule({
  declarations: [DetailLayoutComponent],
  imports: [CommonModule, MatIconModule, BackModule, MatButtonModule, RouterModule],
  exports: [DetailLayoutComponent],
})
export class DetailLayoutModule {}
