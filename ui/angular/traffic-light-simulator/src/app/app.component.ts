import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent implements OnInit {
  title = 'traffic-light-simulator';
  red = true;
  yellow = false;
  green = false;

  ngOnInit() {
    this.trafficSignal();
  }

  trafficSignal() {
    const turnOnRed = () => {
      this.red = true;
      this.yellow = false;
      setTimeout(turnOnYellow, 6500); // Red for 6.5 seconds
    }

    const turnOnYellow = () => {
      this.red = false;
      this.yellow = true;
      setTimeout(turnOnGreen, 1000); // Yellow for 1 second
    }

    const turnOnGreen = () => {
      this.yellow = false;
      this.green = true;
      setTimeout(turnOnYellowAgain, 4000); // Green for 4 seconds
    }

    const turnOnYellowAgain = () => {
      this.green = false;
      this.yellow = true;
      setTimeout(() => { this.trafficSignal() }, 1000); // Yellow again for 1 second, then restart
    }

    // Start the traffic signal
    turnOnRed();
  }
}
