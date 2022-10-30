import { Component, Input, OnInit } from '@angular/core';
import { TIME_ONE_DAY_IN_SECONDS, TIME_ONE_MONTH_IN_SECONDS, TIME_ONE_WEEK_IN_SECONDS } from 'src/app/constant/time';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-job-card',
  templateUrl: './job-card.component.html',
  styleUrls: ['./job-card.component.scss']
})
export class JobCardComponent implements OnInit {

  @Input() topic?: Topic

  filters = [
    { name: "Today", value: { from: Date.now() - TIME_ONE_DAY_IN_SECONDS, to: Date.now() } },
    { name: "Last 7 days", value: { from: Date.now() - TIME_ONE_WEEK_IN_SECONDS, to: Date.now() } },
    { name: "Last 30 days", value: { from: Date.now() - TIME_ONE_MONTH_IN_SECONDS, to: Date.now() } },
  ]
  currentFilter: string = this.filters[0].name

  onClickFilter(filter: string) {
    this.currentFilter = filter
  }

  constructor() { }

  ngOnInit(): void {
  }

}
