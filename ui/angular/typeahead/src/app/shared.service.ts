import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class SharedService {

  constructor(protected  http: HttpClient) { }
  protected getJSON(url: string): Observable<any>{
    return this.http.get(url);
  }
}
