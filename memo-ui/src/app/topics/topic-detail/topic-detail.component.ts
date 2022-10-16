import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';

@Component({
  selector: 'app-topic-detail',
  templateUrl: './topic-detail.component.html',
  styleUrls: ['./topic-detail.component.scss']
})
export class TopicDetailComponent implements OnInit {
  @Input() topic?: Topic;
  @Output() clickClearSelection = new EventEmitter<void>();

  onClickClearSelection() {
    this.clickClearSelection.emit();
  }
  constructor(private topicSelectingService : TopicSelectingService) {
    this.topicSelectingService.selectedTopic.subscribe(topic => {
      this.topic = topic;
    })

   }

  ngOnInit(): void {
  }

}
