import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IndividualEventComponent } from './events/individual-event/individual-event.component';
import { EventFormComponent } from './events/event-form/event-form.component';
import { EventsComponent } from './events/events-list/events.component';

const routes: Routes = [
  { path: 'events', component: EventsComponent },
  { path: 'event/:id', component: IndividualEventComponent },
  { path: 'event/edit/:id', component: EventFormComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
