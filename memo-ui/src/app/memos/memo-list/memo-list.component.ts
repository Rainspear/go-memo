import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-list',
  templateUrl: './memo-list.component.html',
  styleUrls: ['./memo-list.component.scss']
})
export class MemoListComponent implements OnInit {
  @Input() memos: Memo[] = [];
  @Output() clickDeleteMemo = new EventEmitter<string>();

  onClickDeleteMemo(id: string): void {
    this.clickDeleteMemo.emit(id);
  }

  constructor() { }

  ngOnInit(): void {
  }

}
