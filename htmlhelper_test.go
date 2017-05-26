package htmlhelper_test

import (
	"log"

	"github.com/northbright/htmlhelper"
)

func ExampleClean() {
	str = htmlhelper.Clean(str)
	log.Printf("cleaned HTML: %v\n", str)

	// Output:
}

var str = `
</head>

  <body class="logged-in env-production emoji-size-boost page-responsive min-width-0 site-header-dark f4">
    



  <div class="position-relative js-header-wrapper ">
    <a href="#start-of-content" tabindex="1" class="accessibility-aid js-skip-to-content">Skip to content</a>
    <div id="js-pjax-loader-bar" class="pjax-loader-bar"><div class="progress"></div></div>

    
    
    



          <header class="site-header js-details-container Details" role="banner">
  <div class="site-nav-container">
`
