import { Component, Input, OnInit } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-list',
  templateUrl: './memo-list.component.html',
  styleUrls: ['./memo-list.component.scss']
})
export class MemoListComponent implements OnInit {
  @Input() memos: Memo[] = [];
  constructor() { }

  ngOnInit(): void {
  }

}
