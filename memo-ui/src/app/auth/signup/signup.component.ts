import { Input, Output, Component, OnInit, EventEmitter } from '@angular/core';
import { ParamsPostUser } from 'src/app/models/user.model';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {
  @Output() clickCreateUser = new EventEmitter<ParamsPostUser>();
  @Input() error ?: string = "";

  constructor() { }
  user: ParamsPostUser = {email : '', password : '', name : ''};
    ngOnInit(): void {
  }

  onClickCreateUser() { 
    this.clickCreateUser.emit(this.user);
  }
}
