import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';
@Component({
  selector: 'app-topic-item',
  templateUrl: './topic-item.component.html',
  styleUrls: ['./topic-item.component.scss']
})
export class TopicItemComponent implements OnInit {
  @Input() topic?: Topic;
  @Output() selectedTopic = new EventEmitter<Topic>();

  constructor(private topicSelectingService : TopicSelectingService) { }

  onSelectTopic() {
    if (this.topic) this.topicSelectingService.onGetTopic(this.topic);
    // this.selectedTopic.emit(this.topic);
  }

  ngOnInit(): void {
  }

}
