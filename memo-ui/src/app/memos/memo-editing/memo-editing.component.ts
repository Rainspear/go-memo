import { Component, ElementRef, EventEmitter, Input, OnInit, ViewChild,  } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-editing',
  templateUrl: './memo-editing.component.html',
  styleUrls: ['./memo-editing.component.scss']
})
export class MemoEditingComponent implements OnInit {
  clickGetMemoData = new EventEmitter<Memo>();
  @Input() memo?: Memo;
  @ViewChild("TopicTitle") topicTitleInput!: ElementRef;
  @ViewChild("TopicQuestion") topicQuestionInput!: ElementRef;
  @ViewChild("TopicAnswer") topicAnswer!: ElementRef;

  constructor() { }

  onClickAdd(inputElement: HTMLInputElement ) {
    // console.log("current title params pass in function ", inputElement.type + "=" + inputElement.value);
    // console.log("topicTitleInput", this.topicTitleInput);
    // console.log("topicQuestionInput", this.topicQuestionInput);
    // console.log("topicAnswer", this.topicAnswer);
  }

  onClickGetMemoData () {
    this.clickGetMemoData.emit(this.memo);
  }

  ngOnInit(): void {
  }

}
