import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { IndividualEventComponent } from './events/individual-event/individual-event.component';
import { HttpClientModule } from '@angular/common/http';
import { EventFormComponent } from './events/event-form/event-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { FooterComponent } from './footer/footer.component';
import { EventsComponent } from './events/events-list/events.component';
@NgModule({
  declarations: [
    AppComponent,
    EventsComponent,
    IndividualEventComponent,
    EventFormComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
