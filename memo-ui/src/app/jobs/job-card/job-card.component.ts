import { Component, Input, OnInit } from '@angular/core';
import { TIME_ONE_DAY_IN_SECONDS, TIME_ONE_MONTH_IN_SECONDS, TIME_ONE_WEEK_IN_SECONDS } from 'src/app/constant/time';
import { Schedule } from 'src/app/models/schedule.model';
import { IFilterTopic, Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';

@Component({
  selector: 'app-job-card',
  templateUrl: './job-card.component.html',
  styleUrls: ['./job-card.component.scss']
})
export class JobCardComponent implements OnInit {

  @Input() topic?: Topic;
  @Input() schedules?: Schedule[] = [];

  filters: IFilterTopic[] = [
    { name: "Today", value: { from_date: Math.floor(Date.now() / 1000) - TIME_ONE_DAY_IN_SECONDS, to_date: Math.floor(Date.now() / 1000) } },
    { name: "Last 7 days", value: { from_date: Math.floor(Date.now() / 1000) - TIME_ONE_WEEK_IN_SECONDS, to_date: Math.floor(Date.now() / 1000) } },
    { name: "Last 30 days", value: { from_date: Math.floor(Date.now() / 1000) - TIME_ONE_MONTH_IN_SECONDS, to_date: Math.floor(Date.now() / 1000) } },
  ]

  currentFilter: IFilterTopic = this.filters[0];

  onClickFilter(filter: IFilterTopic) {
    this.currentFilter = filter;
    this.topicSelectingService.selectedFilter.next(filter)
  }

  constructor(private topicSelectingService: TopicSelectingService, private apiService: ApiService) {

  }

  ngOnInit(): void {
    this.topicSelectingService.selectedFilter.next(this.currentFilter)
  }

}
