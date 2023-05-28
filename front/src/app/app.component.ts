import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'front';

  constructor(private router: Router) {}

  handleCreateEvent() {
    this.router.navigate(['event/edit/0']).catch(console.error)
  }

  isEditCurrentRoute() {
    return /edit/.test(this.router.url)
  }
}
