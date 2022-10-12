import { Component, ContentChild, ElementRef, OnInit, AfterContentInit, AfterContentChecked } from '@angular/core';

@Component({
  selector: 'app-memo-detail',
  templateUrl: './memo-detail.component.html',
  styleUrls: ['./memo-detail.component.scss']
})
export class MemoDetailComponent implements OnInit, AfterContentInit, AfterContentChecked  {
  @ContentChild("contentChild") contentChild?: ElementRef;
  constructor() { }

  ngOnInit(): void {
    // console.log("ngOnInit contentChild", this.contentChild?.nativeElement)
  }

  ngAfterContentInit(): void {
    // console.log("ngAfterContentInit contentChild", this.contentChild?.nativeElement)
  }

  ngAfterContentChecked(): void {
    // console.log("ngAfterContentChecked contentChild", this.contentChild?.nativeElement) 
  }

}
