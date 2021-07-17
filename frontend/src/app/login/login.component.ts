import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  userForm: FormGroup;

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.userForm = new FormGroup ({
      username: new FormControl(),
      password: new FormControl()
    })
  }

  async registerUser(){
    let httpHeaders = new HttpHeaders({
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin':'*',
      'Cache-Control': 'no-cache'

    });

    let userInfo = this.userForm.value
    console.log(userInfo )


    // this.http
    // .post("http://localhost:8080/register", '{"username": "hey", "password": "boss"}', {
    //   headers: httpHeaders,
    // }).toPromise().then((data: any) => {
    //   console.log(data, 'the****')
    // })

    return this.http
    .post("http://localhost:8080/register", userInfo, {
      headers: httpHeaders,
    })
    .subscribe(
      data => {
        console.log("POST Request is successful ", data);
      },
      error => {
        console.log("Error", error);
      }
    );
  }

}
