function createNode(evento) {
    var e_0 = document.createElement("li");
    e_0.setAttribute("class", "cd-schedule__event");
    var e_1 = document.createElement("a");
    e_1.setAttribute("data-start", evento.start);
    e_1.setAttribute("data-end", evento.end);
    e_1.setAttribute("data-content", "event-yoga-1");
    e_1.setAttribute("data-event", evento.type);
    e_1.setAttribute("href", "#0");
    var e_2 = document.createElement("em");
    e_2.setAttribute("class", "cd-schedule__name");
    e_2.appendChild(document.createTextNode(evento.title));
    e_1.appendChild(e_2);
    e_0.appendChild(e_1);

    placeElement(e_0);
}

function placeElement(node) {

    this.topInfoElement = document.getElementsByClassName('cd-schedule__top-info')[0];
    this.timelineItems = document.getElementsByClassName('cd-schedule__timeline')[0].getElementsByTagName('li');
    this.timelineStart = getScheduleTimestamp(this.timelineItems[0].textContent);
    this.timelineUnitDuration = getScheduleTimestamp(this.timelineItems[1].textContent) - getScheduleTimestamp(this.timelineItems[0].textContent);

    var slotHeight = this.topInfoElement.offsetHeight;
    var anchor = node.getElementsByTagName('a')[0];
    var start = getScheduleTimestamp(anchor.getAttribute('data-start'))
    var duration = getScheduleTimestamp(anchor.getAttribute('data-end')) - start;
    var eventTop = slotHeight*(start - this.timelineStart)/this.timelineUnitDuration
    var eventHeight = slotHeight*duration/this.timelineUnitDuration;

    node.setAttribute('style', 'top: '+(eventTop-1)+'px; height: '+(eventHeight +1)+'px');
    initEvents(node)

    document.getElementById('monday').appendChild(node)
}

function initEvents(node) {

    this.modal = document.getElementsByClassName('cd-schedule-modal')[0];
    this.coverLayer = document.getElementsByClassName('cd-schedule__cover-layer')[0];

    this.modalClose = this.modal.getElementsByClassName('cd-schedule-modal__close')[0];


    // open modal when user selects an event
    node.addEventListener('click', function(event){
        event.preventDefault();
        if(!node.animating) openModal(node.getElementsByTagName('a')[0]);
    });

    //close modal window
    this.modalClose.addEventListener('click', function(event){
        event.preventDefault();
        if( !node.animating ) closeModal(node);
    });

    /*
    this.coverLayer.addEventListener('click', function(event){
        event.preventDefault();
        if( !node.animating ) closeModal(node);
    });
     */
}

function openModal(target) {
    var mq = mqFunc(target);
    this.animating = true;
    this.modalMaxWidth = 800;
    this.modalMaxHeight = 480;

    //update event name and time
    this.modal = document.getElementsByClassName('cd-schedule-modal')[0];
    this.modalEventName = this.modal.getElementsByClassName('cd-schedule-modal__name')[0];
    this.modalEventName.textContent = target.getElementsByTagName('em')[0].textContent;
    this.modalDate = this.modal.getElementsByClassName('cd-schedule-modal__date')[0];
    this.modalDate.textContent = target.getAttribute('data-start')+' - '+target.getAttribute('data-end');
    this.modal.setAttribute('data-event', target.getAttribute('data-event'));

    this.modalHeader = document.getElementsByClassName('cd-schedule-modal__header')[0];
    this.modalHeaderBg = document.getElementsByClassName('cd-schedule-modal__header-bg')[0];
    this.modalBody = document.getElementsByClassName('cd-schedule-modal__body')[0];
    this.modalBodyBg = document.getElementsByClassName('cd-schedule-modal__body-bg')[0];

    //update event content
    this.loadEventContent(target.getAttribute('data-content'));

    Util.addClass(this.modal, 'cd-schedule-modal--open');

    setTimeout(function(){
        //fixes a flash when an event is selected - desktop version only
        Util.addClass(target.closest('li'), 'cd-schedule__event--selected');
    }, 10);

    if( mq == 'mobile' ) {
        this.modal.addEventListener('transitionend', function cb(){
            this.animating = false;
            this.modal.removeEventListener('transitionend', cb);
        });
    } else {
        var eventPosition = target.getBoundingClientRect(),
            eventTop = eventPosition.top,
            eventLeft = eventPosition.left,
            eventHeight = target.offsetHeight,
            eventWidth = target.offsetWidth;

        var windowWidth = window.innerWidth,
            windowHeight = window.innerHeight;

        var modalWidth = ( windowWidth*.8 > this.modalMaxWidth ) ? this.modalMaxWidth : windowWidth*.8,
            modalHeight = ( windowHeight*.8 > this.modalMaxHeight ) ? this.modalMaxHeight : windowHeight*.8;

        var modalTranslateX = parseInt((windowWidth - modalWidth)/2 - eventLeft),
            modalTranslateY = parseInt((windowHeight - modalHeight)/2 - eventTop);

        var HeaderBgScaleY = modalHeight/eventHeight,
            BodyBgScaleX = (modalWidth - eventWidth);

        //change modal height/width and translate it
        this.modal.setAttribute('style', 'top:'+eventTop+'px;left:'+eventLeft+'px;height:'+modalHeight+'px;width:'+modalWidth+'px;transform: translateY('+modalTranslateY+'px) translateX('+modalTranslateX+'px)');
        //set modalHeader width
        this.modalHeader.setAttribute('style', 'width:'+eventWidth+'px');
        //set modalBody left margin
        this.modalBody.setAttribute('style', 'margin-left:'+eventWidth+'px');
        //change modalBodyBg height/width ans scale it
        this.modalBodyBg.setAttribute('style', 'height:'+eventHeight+'px; width: 1px; transform: scaleY('+HeaderBgScaleY+') scaleX('+BodyBgScaleX+')');
        //change modal modalHeaderBg height/width and scale it
        this.modalHeaderBg.setAttribute('style', 'height: '+eventHeight+'px; width: '+eventWidth+'px; transform: scaleY('+HeaderBgScaleY+')');

        this.modalHeaderBg.addEventListener('transitionend', function cb(){
            this.modalHeaderBg = document.getElementsByClassName('cd-schedule-modal__header-bg')[0];
            this.modal = document.getElementsByClassName('cd-schedule-modal')[0];

            //wait for the  end of the modalHeaderBg transformation and show the modal content
            this.animating = false;
            Util.addClass(this.modal, 'cd-schedule-modal--animation-completed');
            this.modalHeaderBg.removeEventListener('transitionend', cb);
        });
    }

    //if browser do not support transitions -> no need to wait for the end of it
    animationFallback();
};

function closeModal(node) {
    var mq = mqFunc(node);
    this.animating = true;
    this.modalMaxWidth = 800;
    this.modalMaxHeight = 480;

    //update event name and time
    this.modal = document.getElementsByClassName('cd-schedule-modal')[0];
    this.modalEventName = this.modal.getElementsByClassName('cd-schedule-modal__name')[0];
    this.modalEventName.textContent = node.getElementsByTagName('em')[0].textContent;
    this.modalDate = this.modal.getElementsByClassName('cd-schedule-modal__date')[0];
    this.modalDate.textContent = node.getAttribute('data-start')+' - '+node.getAttribute('data-end');
    this.modal.setAttribute('data-event', node.getAttribute('data-event'));

    this.modalHeader = document.getElementsByClassName('cd-schedule-modal__header')[0];
    this.modalHeaderBg = document.getElementsByClassName('cd-schedule-modal__header-bg')[0];
    this.modalBody = document.getElementsByClassName('cd-schedule-modal__body')[0];
    this.modalBodyBg = document.getElementsByClassName('cd-schedule-modal__body-bg')[0];

    var mq = mqFunc(node);

    var item = document.getElementsByClassName('cd-schedule__event--selected')[0],
        target = item.getElementsByTagName('a')[0];

    this.animating = true;

    if( mq == 'mobile' ) {
        Util.removeClass(this.modal, 'cd-schedule-modal--open');
        this.modal.addEventListener('transitionend', function cb(){
            Util.removeClass(this.modal, 'cd-schedule-modal--content-loaded');
            Util.removeClass(item, 'cd-schedule__event--selected');
            this.animating = false;
            this.modal.removeEventListener('transitionend', cb);
        });
    } else {
        var eventPosition = target.getBoundingClientRect(),
            eventTop = eventPosition.top,
            eventLeft = eventPosition.left,
            eventHeight = target.offsetHeight,
            eventWidth = target.offsetWidth;

        var modalStyle = window.getComputedStyle(this.modal),
            modalTop = Number(modalStyle.getPropertyValue('top').replace('px', '')),
            modalLeft = Number(modalStyle.getPropertyValue('left').replace('px', ''));

        var modalTranslateX = eventLeft - modalLeft,
            modalTranslateY = eventTop - modalTop;

        Util.removeClass(this.modal, 'cd-schedule-modal--open cd-schedule-modal--animation-completed');

        //change modal width/height and translate it
        this.modal.style.width = eventWidth+'px';this.modal.style.height = eventHeight+'px';this.modal.style.transform = 'translateX('+modalTranslateX+'px) translateY('+modalTranslateY+'px)';
        //scale down modalBodyBg element
        this.modalBodyBg.style.transform = 'scaleX(0) scaleY(1)';
        //scale down modalHeaderBg element
        // self.modalHeaderBg.setAttribute('style', 'transform: scaleY(1)');
        this.modalHeaderBg.style.transform = 'scaleY(1)';

        this.modalHeaderBg.addEventListener('transitionend', function cb(){
            this.modalHeaderBg = document.getElementsByClassName('cd-schedule-modal__header-bg')[0];

            //wait for the  end of the modalHeaderBg transformation and reset modal style
            Util.addClass(this.modal, 'cd-schedule-modal--no-transition');
            setTimeout(function(){
                this.modal.removeAttribute('style');
                this.modalBody.removeAttribute('style');
                this.modalHeader.removeAttribute('style');
                this.modalHeaderBg.removeAttribute('style');
                this.modalBodyBg.removeAttribute('style');
            }, 10);
            setTimeout(function(){
                Util.removeClass(this.modal, 'cd-schedule-modal--no-transition');
            }, 20);
            this.animating = false;
            Util.removeClass(this.modal, 'cd-schedule-modal--content-loaded');
            Util.removeClass(item, 'cd-schedule__event--selected');
            this.modalHeaderBg.removeEventListener('transitionend', cb);
        });
    }

    //if browser do not support transitions -> no need to wait for the end of it
    this.animationFallback();
};


function loadEventContent(content) {
    // load the content of an event when user selects it

    httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                this.modal = document.getElementsByClassName('cd-schedule-modal')[0];

                this.modal.getElementsByClassName('cd-schedule-modal__event-info')[0].innerHTML = getEventContent(httpRequest.responseText);
                Util.addClass(this.modal, 'cd-schedule-modal--content-loaded');
            }
        }
    };
    httpRequest.open('GET', content+'.html');
    httpRequest.send();
};

function getEventContent(string) {
    // reset the loaded event content so that it can be inserted in the modal
    var div = document.createElement('div');
    div.innerHTML = string.trim();
    return div.getElementsByClassName('cd-schedule-modal__event-info')[0].innerHTML;
};
function getScheduleTimestamp(time) {
    //accepts hh:mm format - convert hh:mm to timestamp
    time = time.replace(/ /g,'');
    var timeArray = time.split(':');
    var timeStamp = parseInt(timeArray[0])*60 + parseInt(timeArray[1]);
    return timeStamp;
};

function mqFunc(node){
    //get MQ value ('desktop' or 'mobile')
    return window.getComputedStyle(node, '::before').getPropertyValue('content').replace(/'|"/g, "");
};

function animationFallback() {
    this.modal = document.getElementsByClassName('cd-schedule-modal')[0];
    this.modalHeaderBg = document.getElementsByClassName('cd-schedule-modal__header-bg')[0];

    if( !this.supportAnimation ) { // fallback for browsers not supporting transitions
        var event = new CustomEvent('transitionend');
        this.modal.dispatchEvent(event);
        this.modalHeaderBg.dispatchEvent(event);
    }
};