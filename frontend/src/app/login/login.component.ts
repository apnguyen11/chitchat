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

  loginUser(){
    const headers = {
      'content-type': 'application/json'
    }
    let userInfo = this.userForm.value
    console.log(userInfo )

    return this.http
    .post("api/login", userInfo, {
      headers: headers
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

  registerUser(){
    const headers = {
      'content-type': 'application/json'
    }
    let userInfo = this.userForm.value
    console.log(userInfo )

    return this.http
    .post("api/register", userInfo, {
      headers: headers
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
