import { Component, EventEmitter, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  collapsed: boolean = true;
  @Output() featureSelected = new EventEmitter<string>();
  constructor() { }

  onClickSelectNav(route: string) {
    this.featureSelected.emit(route);
  }

  ngOnInit(): void {
  }

}
