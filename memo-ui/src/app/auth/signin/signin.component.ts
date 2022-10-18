import { Component, Input, OnInit } from '@angular/core';
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

  constructor(private apiService : ApiService, private authService : AuthUserService) { }

  onClickLogin () {
    this.error = "";
    this.apiService.login(this.user).subscribe((res: any) => {
      console.log("res", res)
      if (res.data) {
        console.log(res.data)
        this.authService.onGetUser(res.data)
        return;
      }
      return;
    }, error => {
      this.error = error?.error?.error || error.message
    })
  }

  ngOnInit(): void {
  }

}
