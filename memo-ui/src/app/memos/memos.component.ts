import { Component, Input, OnInit } from '@angular/core';
import { ApiSerivce } from '../services/api.service';
import { Memo } from '../models/memo.model';
import { ResponseAPI } from '../models/response.model';

@Component({
  selector: 'app-memos',
  templateUrl: './memos.component.html',
  styleUrls: ['./memos.component.scss']
})
export class MemosComponent implements OnInit {

  memos: Memo[] = [];
  selectedMemo?: Memo;
  responseStatus: boolean | null = null;
  constructor(private apiSerivce: ApiSerivce) {}

  onClickGetMemoData() {

  }

  onClickDeleteMemo(id: string) {
    this.responseStatus = null;
    this.apiSerivce.deleteMemo(id).subscribe((res : ResponseAPI<Memo>) => {
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
    this.apiSerivce.getAllMemos()
    .subscribe((res: ResponseAPI<Memo[]>) => {
      this.memos = res.data  
    })
  }

}
