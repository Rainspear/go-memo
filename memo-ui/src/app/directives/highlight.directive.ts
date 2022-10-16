import { Directive, ElementRef, HostBinding, HostListener, Input, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective implements OnInit {
  @Input("appHighlight") highlightColor: string = 'red';
  @Input() defaultColor: string = 'black';
  @HostBinding('style.color') color: string = this.defaultColor;

  constructor(private elementRef: ElementRef, private render: Renderer2 ) { }

  ngOnInit(): void {
    this.color = this.defaultColor
    // this.elementRef.nativeElement.style.backgroundColor = 'lightblue';
    // this.elementRef.nativeElement.style.border = '1px solid black';
    // this.render.setStyle(this.elementRef.nativeElement, "background-color", "lightblue");
  }

  @HostListener('mouseenter') onHoverElement(eventData: Event) {
    this.render.setStyle(this.elementRef.nativeElement, "font-size", "2.5rem");
    this.color = this.highlightColor;
  } 

  @HostListener('mouseleave') onLeaveElement(eventData: Event) {
    this.render.setStyle(this.elementRef.nativeElement, "font-size", "inherit");
    this.color = this.defaultColor;
  } 

}
