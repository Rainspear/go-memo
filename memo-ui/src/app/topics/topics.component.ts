import { Component, OnInit } from '@angular/core';
import { ApiSerivce } from '../api.service';
import { Topic } from '../models/topic.model';

@Component({
  selector: 'app-topics',
  templateUrl: './topics.component.html',
  styleUrls: ['./topics.component.scss']
})
export class TopicsComponent implements OnInit {

  topics ?: Topic[];

  constructor(private apiSerivce : ApiSerivce) { }

  ngOnInit(): void {
    this.apiSerivce.getAllTopics()
     .subscribe((res: any) => {
      this.topics = res.data
     })
  }
}
