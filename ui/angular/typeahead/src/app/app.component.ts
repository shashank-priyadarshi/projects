import { Component } from '@angular/core';
import {SharedService} from "./shared.service";
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent extends SharedService{
  private typingTimer: any;
  title = 'typeahead';

  constructor(http: HttpClient) {
    super(http);
  }

  private debounce(callback: Function, delay:number){
    clearTimeout(this.typingTimer);
    this.typingTimer = setTimeout(() =>{callback()}, delay);
  }

  onInputChange(input: HTMLInputElement) {
    const inputValue = input.value;

    this.debounce(() => {
      console.log('Function called after typing stops:', inputValue);
      // Call your desired function here with inputValue
      console.log(this.getJSON("assets/data/"+input.id+".json").subscribe((obj)=>{
        console.log("subscription object: ",obj);
      }));
    }, 1000); // 1000ms (1 second) debounce delay
  }

//   flatten the array
//   every property of an array will be prefixed by the id/name
//   save this value to a map, and allow searching
//   if value if already available in localStorage, just search from localStorage
//   for the value selected, display the whole object, provide option to see the flattened object
//   TODO: allow searching based on both keys and values
}
