import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { faPlusCircle } from '@fortawesome/free-solid-svg-icons';
import { Memo } from 'src/app/models/memo.model';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-memo-list',
  templateUrl: './memo-list.component.html',
  styleUrls: ['./memo-list.component.scss']
})
export class MemoListComponent implements OnInit {
  faPlusCircle = faPlusCircle;

  @Input() memos?: Memo[] ;
  @Input() topic?: Topic;
  @Output() clickDeleteMemo = new EventEmitter<string>();
  showCreatingMemo = false;
  onClickDeleteMemo(id: string): void {
    this.clickDeleteMemo.emit(id);
  }

  onToggle(show: boolean){
    this.showCreatingMemo = show;
  }

  constructor() { }

  ngOnInit(): void {
  }

}
