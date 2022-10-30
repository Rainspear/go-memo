import { Component, OnInit } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Topic } from '../models/topic.model';
import { ResponseAPI } from '../models/response.model';
import { TopicSelectingService } from '../services/topic-selecting.service';

@Component({
  selector: 'app-topics',
  templateUrl: './topics.component.html',
  styleUrls: ['./topics.component.scss']
})
export class TopicsComponent implements OnInit {
  selectedTopic?: Topic;
  topics: Topic[] = [];

  constructor(private apiSerivce : ApiService, private topicSelectingService : TopicSelectingService) {
    this.topicSelectingService.selectedTopic.subscribe((topic) => {
      if (topic) this.invokeAllTopic()
    })
   }

  onClickSelectTopic(topic : Topic) {
    console.log("topic", topic);
    this.selectedTopic = topic;
  }

  onClickClearSelection(){
    console.log("Clear selection");
    this.selectedTopic = undefined;
  }

  invokeAllTopic() {
    this.apiSerivce.getAllTopics()
     .subscribe((res: ResponseAPI<Topic[]>) => {
      this.topics = res.data
     })
  }

  ngOnInit(): void {
    this.invokeAllTopic();
  }
}
