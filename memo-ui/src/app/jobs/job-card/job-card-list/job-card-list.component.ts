import { Component, Input, OnInit } from '@angular/core';
import { Repetition } from 'src/app/models/topic.model';

@Component({
  selector: 'app-job-card-list',
  templateUrl: './job-card-list.component.html',
  styleUrls: ['./job-card-list.component.scss']
})
export class JobCardListComponent implements OnInit {

  @Input() jobs?: Repetition[]
  constructor() { }

  ngOnInit(): void {
  }

}
