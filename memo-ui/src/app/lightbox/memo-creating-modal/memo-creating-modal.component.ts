import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Memo, ParamsCreateMemo } from 'src/app/models/memo.model';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-memo-creating-modal',
  templateUrl: './memo-creating-modal.component.html',
  styleUrls: ['./memo-creating-modal.component.scss']
})
export class MemoCreatingModalComponent implements OnInit {

  memo: ParamsCreateMemo = { content: '', question: '', answer: [], topic_id :'' };
  @Input() show: boolean = false;
  @Output() toggle = new EventEmitter<boolean>();
  @Output() submit = new EventEmitter<void>();
  error?: string;

  onToggleHandler(show: boolean) {
    this.toggle.emit(show);
  }

  onSubmitHandler() {
    this.error = undefined;
    if (this.memo.question && this.memo.content && this.memo.topic_id) {
      this.apiService.createMemo(this.memo).subscribe((res: any) => {
        if (res.data) {
          this.toggle.emit(false)
          this.submit.emit()
        }
      }, (error) => {
        this.error = error.error
      })
    }
  }

  constructor(private apiService: ApiService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.memo.topic_id = this.route.snapshot.params['id'];
  }

}
