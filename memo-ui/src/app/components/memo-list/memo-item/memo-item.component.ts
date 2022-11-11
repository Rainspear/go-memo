import { Component, EventEmitter, Input, OnInit, Output, AfterViewChecked } from '@angular/core';
import { faCircleQuestion, faNoteSticky, faXmark } from '@fortawesome/free-solid-svg-icons';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-item',
  templateUrl: './memo-item.component.html',
  styleUrls: ['./memo-item.component.scss']
})
export class MemoItemComponent implements OnInit, AfterViewChecked {
  faCircleQuestion = faCircleQuestion;
  faNoteSticky = faNoteSticky;
  faXmark = faXmark;


  @Output() clickDeleteMemo = new EventEmitter<string>();
  @Input() memo?: Memo;
  @Input() index: number = 0;

  constructor() { }


  onClickDeleteMemo() {
    this.clickDeleteMemo.emit(this.memo?.id);
  }

  onSelect() {
    // if (this.memo) this.memoDetailService.selectMemo.next(this.memo);
  }


  ngOnInit(): void {
  }

  ngAfterViewChecked(): void {
  }
}
