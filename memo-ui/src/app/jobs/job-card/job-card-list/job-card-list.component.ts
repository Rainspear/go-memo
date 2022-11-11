import { Component, Input, OnInit, OnChanges, SimpleChanges } from '@angular/core';
import { Schedule } from 'src/app/models/schedule.model';

@Component({
  selector: 'app-job-card-list',
  templateUrl: './job-card-list.component.html',
  styleUrls: ['./job-card-list.component.scss']
})
export class JobCardListComponent implements OnInit, OnChanges {

  @Input() schedules?: Schedule[] = undefined
  constructor() { }

  ngOnChanges(changes: any): void {
    console.log("changes ", changes)
  }

  ngOnInit(): void {
  }

}
