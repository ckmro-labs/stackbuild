import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { BooksComponent } from './books.component';
import { BooksRoutingModule } from './books-routing.module'
@NgModule({
  declarations: [BooksComponent],
  imports: [
    BooksRoutingModule,
    NgbModule.forRoot()
  ],
  exports: [BooksComponent]
})
export class BooksModule { }
