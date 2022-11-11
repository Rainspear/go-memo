import {Input, Output, EventEmitter, Component, OnInit } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-detail-modal',
  templateUrl: './memo-detail-modal.component.html',
  styleUrls: ['./memo-detail-modal.component.scss']
})
export class MemoDetailModalComponent implements OnInit {
  @Input() show: boolean = false;
  @Output() toggle = new EventEmitter<boolean>()
  error?: string;
  @Input() memo?: Memo;
  
  onToggleHandler(show: boolean) {
    this.toggle.emit(show);
  }

  onSubmitHandler() {

  }

  constructor() { }

  ngOnInit(): void {
  }

}
