import { Component, OnInit } from '@angular/core';
import { ParamsPostUser, User } from '../models/user.model';
import { ApiSerivce } from '../services/api.service';
@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss'],
  providers: [ApiSerivce]
})
export class AuthComponent implements OnInit {

  route: string = 'signup';
  user?: User
  error?: string = ""
  constructor(private apiSerivce: ApiSerivce) { }

  onClickCreateUser(user: ParamsPostUser) {
    this.createUser(user);
  }

  createUser(user: ParamsPostUser): void {
    this.error = "";
    this.apiSerivce.createUser(user).subscribe((res: any) => {
      console.log("res", res)
      if (res.data) {
        console.log(res.data)
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
