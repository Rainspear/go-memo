import { Component, Input, OnInit, AfterViewChecked } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Memo } from '../models/memo.model';
import { ResponseAPI } from '../models/response.model';
import { Topic } from '../models/topic.model';

@Component({
  selector: 'app-memos',
  templateUrl: './memos.component.html',
  styleUrls: ['./memos.component.scss']
})
export class MemosComponent implements OnInit, AfterViewChecked {
  @Input() memos: Memo[] = [];
  selectedMemo?: Memo;
  responseStatus: boolean | null = null;
  constructor(private apiSerivce: ApiService) { }

  // onClickGetMemoData() {

  // }

  ngAfterViewChecked(): void {
  }

  onClickDeleteMemo(id: string) {
    this.responseStatus = null;
    this.apiSerivce.deleteMemo(id).subscribe((res: ResponseAPI<Memo>) => {
      if (res?.data) {
        this.responseStatus = true;
        this.apiSerivce.getAllMemos()
          .subscribe((res: ResponseAPI<Memo[]>) => {
            this.memos = res.data;
          })
        return
      }
      this.responseStatus = true;
      return false;
    });
  }

  ngOnInit(): void {
    // this.apiSerivce.getAllMemos()
    //   .subscribe((res: ResponseAPI<Memo[]>) => {
    //     this.memos = res.data
    //   })
  }

}
