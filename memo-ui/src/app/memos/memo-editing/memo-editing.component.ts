import { Component, EventEmitter, Input, OnInit } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-editing',
  templateUrl: './memo-editing.component.html',
  styleUrls: ['./memo-editing.component.scss']
})
export class MemoEditingComponent implements OnInit {
  clickGetMemoData = new EventEmitter<Memo>();
  @Input() memo?: Memo;
  constructor() { }

  onClickGetMemoData () {
    this.clickGetMemoData.emit(this.memo);
  }

  ngOnInit(): void {
  }

}
