import { Component, ContentChild, ElementRef, OnInit, AfterContentInit, AfterContentChecked, Input } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-detail',
  templateUrl: './memo-detail.component.html',
  styleUrls: ['./memo-detail.component.scss']
})
export class MemoDetailComponent implements OnInit, AfterContentInit, AfterContentChecked  {
  @ContentChild("contentChild") contentChild?: ElementRef;
  @Input() memo?: Memo;
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
