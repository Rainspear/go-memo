import { Component, Input, OnInit } from '@angular/core';
import { Schedule } from 'src/app/models/schedule.model';

@Component({
  selector: 'app-job-card-list',
  templateUrl: './job-card-list.component.html',
  styleUrls: ['./job-card-list.component.scss']
})
export class JobCardListComponent implements OnInit {

  @Input() jobs?: Schedule[] = []
  constructor() { }

  ngOnInit(): void {
  }

}
