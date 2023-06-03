import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Event } from '../models/event';

@Component({
  selector: 'app-individual-event',
  templateUrl: './individual-event.component.html',
  styleUrls: ['./individual-event.component.scss'],
})
export class IndividualEventComponent {
  public event!: Event;

  constructor(private route: ActivatedRoute, private http: HttpClient) {}

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');

    this.http
      .get<Event>(`http://0.0.0.0:8080/events/${id}`, { observe: 'body' })
      .subscribe((val: Event) => {
        this.event = val;

        console.log(this.event);
      });
  }
}
