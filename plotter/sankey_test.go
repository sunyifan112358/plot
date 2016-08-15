// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"image/color"
	"log"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/vg/draw"
)

// ExampleSankey_sample creates a simple sankey diagram.
// The output can be found at https://github.com/gonum/plot/blob/master/plotter/testdata/sankeySimple_golden.png.
func ExampleSankey_simple() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	// Define the stock categories
	const (
		treeType int = iota
		consumer
		fate
	)
	categoryLabels := []string{"Tree type", "Consumer", "Fate"}

	flows := []*Flow{
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Mohamed",
			Value:                 5,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Mohamed",
			Value:                 2,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Sofia",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Sofia",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Wei",
			Value:                 6,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Value:                 6,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Sofia",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Sofia",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Value:                 0.5, // An unbalanced flow
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Wei",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Value:                 5,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Wei",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Value:                 0.3,
		},
	}

	sankey, err := NewSankey(flows...)
	if err != nil {
		log.Panic(err)
	}
	p.Add(sankey)
	p.Y.Label.Text = "Number of apples"
	p.NominalX(categoryLabels...)
	err = p.Save(300, 180, "testdata/sankeySimple.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestSankey_simple(t *testing.T) {
	checkPlot(ExampleSankey_simple, t, "sankeySimple.png")
}

// ExampleSankey_grouped creates a sankey diagram with grouped flows.
// The output can be found at https://github.com/gonum/plot/blob/master/plotter/testdata/sankeyGrouped_golden.png.
func ExampleSankey_grouped() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	// Define the stock categories
	const (
		treeType int = iota
		consumer
		fate
	)
	categoryLabels := []string{"Tree type", "Consumer", "Fate"}

	flows := []*Flow{
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Mohamed",
			Group:                 "Apples",
			Value:                 5,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Mohamed",
			Group:                 "Dates",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Mohamed",
			Group:                 "Lychees",
			Value:                 2,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Sofia",
			Group:                 "Apples",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Sofia",
			Group:                 "Dates",
			Value:                 4,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Sofia",
			Group:                 "Apples",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Wei",
			Group:                 "Lychees",
			Value:                 6,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: consumer,
			ReceptorStockLabel:    "Wei",
			Group:                 "Apples",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Apples",
			Value:                 4,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Apples",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Dates",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Mohamed",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Lychees",
			Value:                 2,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Sofia",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Apples",
			Value:                 4,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Sofia",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Dates",
			Value:                 3,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Sofia",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Dates",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Wei",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Lychees",
			Value:                 6,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Wei",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Eaten",
			Group:                 "Apples",
			Value:                 2,
		},
		&Flow{
			SourceStockCategory:   consumer,
			SourceStockLabel:      "Wei",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Apples",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Apples",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Large",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Dates",
			Value:                 1,
		},
		&Flow{
			SourceStockCategory:   treeType,
			SourceStockLabel:      "Small",
			ReceptorStockCategory: fate,
			ReceptorStockLabel:    "Waste",
			Group:                 "Lychees",
			Value:                 0.3,
		},
	}

	sankey, err := NewSankey(flows...)
	if err != nil {
		log.Panic(err)
	}

	// Here we specify the FLowStyle function to set the
	// colors of the different fruit groups.
	sankey.FlowStyle = func(group string) (color.Color, draw.LineStyle) {
		switch group {
		case "Lychees":
			return color.NRGBA{R: 242, G: 169, B: 178, A: 100}, sankey.LineStyle
		case "Apples":
			return color.NRGBA{R: 91, G: 194, B: 54, A: 100}, sankey.LineStyle
		case "Dates":
			return color.NRGBA{R: 112, G: 22, B: 0, A: 100}, sankey.LineStyle
		default:
			panic(fmt.Errorf("invalid group %s", group))
		}
	}

	// Here we set the backgroud color for stocks from grey to white.
	sankey.Color = color.White

	p.Add(sankey)
	p.Y.Label.Text = "Number of fruit pieces"
	p.NominalX(categoryLabels...)

	legendLabels, thumbs := sankey.Thumbnailers()
	for i, l := range legendLabels {
		t := thumbs[i]
		p.Legend.Add(l, t)
	}
	p.Legend.Top = true
	p.X.Max = 3.05 // give room for the legend

	err = p.Save(300, 180, "testdata/sankeyGrouped.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestSankey_grouped(t *testing.T) {
	checkPlot(ExampleSankey_grouped, t, "sankeyGrouped.png")
}
