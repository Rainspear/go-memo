import { Component, EventEmitter, Input, OnInit, Output, AfterViewChecked } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';
import { MemoDetailService } from 'src/app/services/memo-detail.service';

@Component({
  selector: 'app-memo-item',
  templateUrl: './memo-item.component.html',
  styleUrls: ['./memo-item.component.scss']
})
export class MemoItemComponent implements OnInit, AfterViewChecked {
  @Output() clickDeleteMemo = new EventEmitter<string>();
  @Input() memo?: Memo;
  @Input() index: number = 0;

  onClickDeleteMemo() {
    this.clickDeleteMemo.emit(this.memo?.id);
  }

  onSelect() {
    if (this.memo) this.memoDetailService.selectMemo.next(this.memo);
  }

  constructor(private memoDetailService: MemoDetailService) { }

  ngOnInit(): void {
  }

  ngAfterViewChecked(): void {
  }
}
