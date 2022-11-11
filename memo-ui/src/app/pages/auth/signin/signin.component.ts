import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ParamsLoginUser } from 'src/app/models/user.model';
import { ApiService } from 'src/app/services/api.service';
import { AuthUserService } from 'src/app/services/auth-user.service';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.scss']
})
export class SigninComponent implements OnInit {
  user: ParamsLoginUser = {email: '', password: ''};

  @Input() error ?: string = "";

  constructor(private apiService : ApiService, private authService : AuthUserService, private router: Router) { }

  onClickLogin () {
    this.error = "";
    this.apiService.login(this.user).subscribe((res: any) => {
      if (res.token) {
        localStorage.setItem('token', res.token)
        this.router.navigate(['/topic']);
      }
    }, error => {
      this.error = error?.error?.error || error.message
    })
  }

  ngOnInit(): void {
  }

}
