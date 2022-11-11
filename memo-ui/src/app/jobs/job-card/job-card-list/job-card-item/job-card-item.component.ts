import { Component, Input, OnInit } from '@angular/core';
import { Schedule } from 'src/app/models/schedule.model';

@Component({
  selector: 'app-job-card-item',
  templateUrl: './job-card-item.component.html',
  styleUrls: ['./job-card-item.component.scss']
})
export class JobCardItemComponent implements OnInit {
  @Input() index: number = 0;
  @Input() job?: Schedule
  constructor() { }

  ngOnInit(): void {
  }

}
