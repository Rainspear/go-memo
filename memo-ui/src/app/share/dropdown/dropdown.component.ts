import { Component, ElementRef, EventEmitter, HostListener, Input, OnInit, Output, ViewChild } from '@angular/core';
import { DropdownProps } from 'src/app/constant/dropdown';

@Component({
  selector: 'app-dropdown',
  templateUrl: './dropdown.component.html',
  styleUrls: ['./dropdown.component.scss']
})
export class DropdownComponent implements OnInit {
  @Input() value?: DropdownProps; // current selected value of dropdown
  @Input() options: DropdownProps[] = [];
  @Output() select = new EventEmitter<DropdownProps>();
  @ViewChild('toggleDropdownButton') toggleDropdownButton!: ElementRef;
  @ViewChild('toggleDropdownButton') menuDropdown!: ElementRef;
  
  showDropdown: boolean = false; // for click outside and turn off menu

  ngOnInit(): void {
  }

  onSelectOption(opt: DropdownProps): void {
    this.value = opt;
    this.select.emit(this.value);
  }

  onToggleDropdown(): void {
    this.showDropdown = !this.showDropdown;
  }

  @HostListener("document:click", ['$event'])
  clickedOutside(e: MouseEvent) {
    e.preventDefault();
    e.stopPropagation();
    if (e.target !== this.toggleDropdownButton.nativeElement && e.target !== this.menuDropdown.nativeElement) {
      // this.menuDropdown = false;
      this.showDropdown = false;
    }
  }
}
