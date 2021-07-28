import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { map, takeUntil, tap } from 'rxjs/operators';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  loginError = false;
  userForm: FormGroup;

  constructor(
    private http: HttpClient,
    private router: Router
    ) { }

  ngOnInit(): void {
    this.userForm = new FormGroup ({
      username: new FormControl(),
      password: new FormControl()
    })

    this.getSession()
  }

  getSession(){
    var session = this.http.get("/").pipe(map((res: any)=>res.json()));
    console.log(session.subscribe((posts)=>{
      console.log("get session",posts);
    }));
  }

  loginUser(){
    const headers = {
      'content-type': 'application/json'
    }
    let userInfo = this.userForm.value
    console.log(userInfo )

    this.getSession()

    return this.http
    .post("api/login", userInfo, {
      headers: headers,
      withCredentials: true
    })
    .subscribe(
      (data: any) => {
        if(data.success){
          this.router.navigate(['/chat']);
        } else {
          this.loginError = true;
        }
        console.log("POST Request is successful ", data.success);
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
