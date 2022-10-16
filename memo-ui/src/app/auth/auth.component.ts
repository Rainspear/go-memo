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
  statusAction: boolean | null = null;
  constructor(private apiSerivce: ApiSerivce) { }

  onClickCreateUser(user: ParamsPostUser) {
    this.createUser(user);
  }

  createUser(user: ParamsPostUser): void {
    this.statusAction = false;
    this.apiSerivce.createUser(user).subscribe((res: any) => {
      if (res.data) {
        console.log(res.data)
        this.statusAction = true;
        return;
      }
      this.statusAction = false;
      return;
    })
    console.log("After created", this.statusAction)
  }

  ngOnInit(): void {
  }

}
