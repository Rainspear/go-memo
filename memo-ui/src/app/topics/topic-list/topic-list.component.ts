import { Component, EventEmitter, Input, OnInit, Output, SimpleChanges } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-topic-list',
  templateUrl: './topic-list.component.html',
  styleUrls: ['./topic-list.component.scss']
})
export class TopicListComponent implements OnInit {
  @Input() topics: Topic[] = [];
  @Output() selectedTopic = new EventEmitter<Topic>();

  onSelectTopic(topic : Topic): void {
    this.selectedTopic.emit(topic);
  }

  ngOnInit(): void {
    console.log("this.topics", this.topics)
  }

  ngOnChanges(changes: SimpleChanges) {
    // console.log("TopicListComponent ngOnChanges: ", changes)
    // this.topics = changes['topics']
  }

}
