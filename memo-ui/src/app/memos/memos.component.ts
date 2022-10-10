import { Component, Input, OnInit } from '@angular/core';
import { ApiSerivce } from '../api.service';
import { Memo } from '../models/memo.model';

@Component({
  selector: 'app-memos',
  templateUrl: './memos.component.html',
  styleUrls: ['./memos.component.scss']
})
export class MemosComponent implements OnInit {

  memos: Memo[] = [];

  constructor(private apiSerivce: ApiSerivce) {}

  ngOnInit(): void {
    this.apiSerivce.getAllMemos()
    .subscribe((res: any) => {
      this.memos = res.data  
    })
  }

}
