import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { faCaretLeft, faGear } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/models/user.model';
import { AuthUserService } from 'src/app/services/auth-user.service';

@Component({
  selector: 'app-share-header',
  templateUrl: './share-header.component.html',
  styleUrls: ['./share-header.component.scss']
})
export class ShareHeaderComponent implements OnInit {
  faGear = faGear;
  faCaretLeft = faCaretLeft;

  user?: User

  constructor(private authUserService: AuthUserService) { }

  ngOnInit(): void {
    this.user = this.authUserService.user  
  }
}
