import { Component, OnInit } from '@angular/core';
import { ApiSerivce } from '../services/api.service';
import { Topic } from '../models/topic.model';
import { ResponseAPI } from '../models/response.model';

@Component({
  selector: 'app-topics',
  templateUrl: './topics.component.html',
  styleUrls: ['./topics.component.scss']
})
export class TopicsComponent implements OnInit {
  selectedTopic?: Topic;
  topics: Topic[] = [];

  constructor(private apiSerivce : ApiSerivce) { }

  onClickSelectTopic(topic : Topic) {
    console.log("topic", topic);
    this.selectedTopic = topic;
  }

  ngOnInit(): void {
    this.apiSerivce.getAllTopics()
     .subscribe((res: ResponseAPI<Topic[]>) => {
      this.topics = res.data
     })
  }
}
