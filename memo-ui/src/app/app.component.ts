import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'memo-ui';
  selectedRoute = "topic"
  onNavigate (route: string) {
    this.selectedRoute = route;
    console.log("this.selectedRoute", this.selectedRoute)
  }
}
