import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private http: HttpClient) { }
  title = 'frontend';
  message;
  inputMessage: any = {"channel": "#ANGULAR", "username": "andy", "content": "ANGULAR MESSAGE"};

  ngOnInit() {
    // Simple GET request with response type <any>


      this.http.get('http://localhost:8080/messages/receive', {responseType: 'text'}).subscribe(data => {
        this.message = data;
      })

  }


  onKey(event: any) { // without type info
    this.inputMessage = event.target.value;
    console.log(this.inputMessage, 'var')
    console.log(event, 'event')
  }

  sendMessage(): any{
    console.log('send btn clicked', this.inputMessage);
    const headers = { 'content-type': 'application/json'}
    var data = {'channel': '#shouting', 'username': 'me', 'content': this.inputMessage}
    this.http.post('http://localhost:8080/messages/send', data, {headers: headers}).subscribe();
  }
}
