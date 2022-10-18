import { Input, Output, Component, OnInit, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { ParamsCreateUser } from 'src/app/models/user.model';
import { ApiService } from 'src/app/services/api.service';
import { AuthUserService } from 'src/app/services/auth-user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {
  // @Output() clickCreateUser = new EventEmitter<ParamsCreateUser>();
  @Input() error ?: string = "";

  constructor(private apiSerivce: ApiService, private authService :AuthUserService,private router: Router) { }
  user: ParamsCreateUser = {email : '', password : '', name : ''};
    ngOnInit(): void {
  }

  onNavigate() {
    this.router.navigate(["/auth","signin"])
  }

  onClickCreateUser() {
    this.createUser(this.user);
  }

  createUser(user: ParamsCreateUser): void {
    this.error = "";
    this.apiSerivce.createUser(user).subscribe((res: any) => {
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

  // onClickCreateUser() { 
  //   this.clickCreateUser.emit(this.user);
  // }
}
