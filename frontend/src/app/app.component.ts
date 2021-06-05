import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private http: HttpClient) { }
  title = 'frontend';
  message;

  ngOnInit() {
    // Simple GET request with response type <any>
    this.http.get('http://localhost:8080/messages/receive', {responseType: 'text'}).subscribe(data => {
        this.message = data;
        console.log("test");
        console.log(this.message);
    })
}
}
