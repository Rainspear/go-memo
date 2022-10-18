import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ParamsCreateUser, User } from '../models/user.model';
import { ApiService } from '../services/api.service';
import { AuthUserService } from '../services/auth-user.service';
@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss'],
  providers: [ApiService]
})

export class AuthComponent implements OnInit {

  route: string = 'signup';
  user?: User
  error?: string = ""
  constructor(private apiService : ApiService, private authService: AuthUserService) { }


  ngOnInit(): void {

  }

}
