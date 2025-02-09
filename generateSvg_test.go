package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSvg(t *testing.T) {
	t.Run("Generate empty", func(t *testing.T) {
		languages := []*Language{}
		res := generateSvg(languages, "testEmpty")
		assert.Equal(t, "\n<svg width=\"300\" height=\"95\" viewBox=\"0 0 300 95\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" role=\"img\" aria-labelledby=\"descId\">\n    <style>\n        .header {\n            font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif;\n            fill: #41b883;\n            animation: fadeInAnimation 0.8s ease-in-out forwards;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .header { font-size: 15.5px; }\n        }\n\n        @keyframes slideInAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: calc(100%-100px);\n            }\n        }\n        @keyframes growWidthAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: 100%;\n            }\n        }\n        .stat {\n            font: 600 14px 'Segoe UI', Ubuntu, \"Helvetica Neue\", Sans-Serif; fill: #fffefe;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .stat { font-size:12px; }\n        }\n        .bold { font-weight: 700 }\n        .lang-name {\n            font: 400 11px \"Segoe UI\", Ubuntu, Sans-Serif;\n            fill: #fffefe;\n        }\n        .stagger {\n            opacity: 0;\n            animation: fadeInAnimation 0.3s ease-in-out forwards;\n        }\n        #rect-mask rect{\n            animation: slideInAnimation 1s ease-in-out forwards;\n        }\n        .lang-progress{\n            animation: growWidthAnimation 0.6s ease-in-out forwards;\n        }\n\n\n\n         \n        @keyframes scaleInAnimation {\n            from {\n                transform: translate(-5px, 5px) scale(0);\n            }\n            to {\n                transform: translate(-5px, 5px) scale(1);\n            }\n        }\n        @keyframes fadeInAnimation {\n            from {\n                opacity: 0;\n            }\n            to {\n                opacity: 1;\n            }\n        }\n    </style>\n    <rect x=\"0.5\" y=\"0.5\" rx=\"4.5\" height=\"99%\" stroke=\"#e4e2e2\" width=\"299\" fill=\"#273849\" stroke-opacity=\"0\"/>\n    <g transform=\"translate(25, 35)\">\n        <g transform=\"translate(0, 0)\">\n            <text x=\"0\" y=\"0\" class=\"header\">Most Used Languages</text>\n        </g>\n    </g>\n\n    <g transform=\"translate(0, 55)\">\n        <svg x=\"25\">\n            <mask id=\"rect-mask\">\n                <rect x=\"0\" y=\"0\" width=\"250\" height=\"8\" fill=\"white\" rx=\"5\"/>\n            </mask>\n            \n\n            <g transform=\"translate(0, 25)\">\n                <g transform=\"translate(0, 0)\">\n                    \n                </g>\n                <g transform=\"translate(150, 0)\">\n                    \n                </g>\n            </g>\n        </svg>\n    </g>\n</svg>\n", res)
	})
	t.Run("Generate empty output", func(t *testing.T) {
		languages := []*Language{}
		res := generateSvg(languages, "")
		assert.Equal(t, "\n<svg width=\"300\" height=\"95\" viewBox=\"0 0 300 95\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" role=\"img\" aria-labelledby=\"descId\">\n    <style>\n        .header {\n            font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif;\n            fill: #41b883;\n            animation: fadeInAnimation 0.8s ease-in-out forwards;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .header { font-size: 15.5px; }\n        }\n\n        @keyframes slideInAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: calc(100%-100px);\n            }\n        }\n        @keyframes growWidthAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: 100%;\n            }\n        }\n        .stat {\n            font: 600 14px 'Segoe UI', Ubuntu, \"Helvetica Neue\", Sans-Serif; fill: #fffefe;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .stat { font-size:12px; }\n        }\n        .bold { font-weight: 700 }\n        .lang-name {\n            font: 400 11px \"Segoe UI\", Ubuntu, Sans-Serif;\n            fill: #fffefe;\n        }\n        .stagger {\n            opacity: 0;\n            animation: fadeInAnimation 0.3s ease-in-out forwards;\n        }\n        #rect-mask rect{\n            animation: slideInAnimation 1s ease-in-out forwards;\n        }\n        .lang-progress{\n            animation: growWidthAnimation 0.6s ease-in-out forwards;\n        }\n\n\n\n         \n        @keyframes scaleInAnimation {\n            from {\n                transform: translate(-5px, 5px) scale(0);\n            }\n            to {\n                transform: translate(-5px, 5px) scale(1);\n            }\n        }\n        @keyframes fadeInAnimation {\n            from {\n                opacity: 0;\n            }\n            to {\n                opacity: 1;\n            }\n        }\n    </style>\n    <rect x=\"0.5\" y=\"0.5\" rx=\"4.5\" height=\"99%\" stroke=\"#e4e2e2\" width=\"299\" fill=\"#273849\" stroke-opacity=\"0\"/>\n    <g transform=\"translate(25, 35)\">\n        <g transform=\"translate(0, 0)\">\n            <text x=\"0\" y=\"0\" class=\"header\">Most Used Languages</text>\n        </g>\n    </g>\n\n    <g transform=\"translate(0, 55)\">\n        <svg x=\"25\">\n            <mask id=\"rect-mask\">\n                <rect x=\"0\" y=\"0\" width=\"250\" height=\"8\" fill=\"white\" rx=\"5\"/>\n            </mask>\n            \n\n            <g transform=\"translate(0, 25)\">\n                <g transform=\"translate(0, 0)\">\n                    \n                </g>\n                <g transform=\"translate(150, 0)\">\n                    \n                </g>\n            </g>\n        </svg>\n    </g>\n</svg>\n", res)
	})

	t.Run("Generate with languages", func(t *testing.T) {
		languages := []*Language{
			{Name: "Go", Color: "#00ADD8", Size: 100, Percentage: 50},
			{Name: "PHP", Color: "#4F5D95", Size: 100, Percentage: 50},
		}
		res := generateSvg(languages, "testEmpty")
		assert.Equal(t, "\n<svg width=\"300\" height=\"120\" viewBox=\"0 0 300 120\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" role=\"img\" aria-labelledby=\"descId\">\n    <style>\n        .header {\n            font: 600 18px 'Segoe UI', Ubuntu, Sans-Serif;\n            fill: #41b883;\n            animation: fadeInAnimation 0.8s ease-in-out forwards;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .header { font-size: 15.5px; }\n        }\n\n        @keyframes slideInAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: calc(100%-100px);\n            }\n        }\n        @keyframes growWidthAnimation {\n            from {\n                width: 0;\n            }\n            to {\n                width: 100%;\n            }\n        }\n        .stat {\n            font: 600 14px 'Segoe UI', Ubuntu, \"Helvetica Neue\", Sans-Serif; fill: #fffefe;\n        }\n        @supports(-moz-appearance: auto) {\n             \n            .stat { font-size:12px; }\n        }\n        .bold { font-weight: 700 }\n        .lang-name {\n            font: 400 11px \"Segoe UI\", Ubuntu, Sans-Serif;\n            fill: #fffefe;\n        }\n        .stagger {\n            opacity: 0;\n            animation: fadeInAnimation 0.3s ease-in-out forwards;\n        }\n        #rect-mask rect{\n            animation: slideInAnimation 1s ease-in-out forwards;\n        }\n        .lang-progress{\n            animation: growWidthAnimation 0.6s ease-in-out forwards;\n        }\n\n\n\n         \n        @keyframes scaleInAnimation {\n            from {\n                transform: translate(-5px, 5px) scale(0);\n            }\n            to {\n                transform: translate(-5px, 5px) scale(1);\n            }\n        }\n        @keyframes fadeInAnimation {\n            from {\n                opacity: 0;\n            }\n            to {\n                opacity: 1;\n            }\n        }\n    </style>\n    <rect x=\"0.5\" y=\"0.5\" rx=\"4.5\" height=\"99%\" stroke=\"#e4e2e2\" width=\"299\" fill=\"#273849\" stroke-opacity=\"0\"/>\n    <g transform=\"translate(25, 35)\">\n        <g transform=\"translate(0, 0)\">\n            <text x=\"0\" y=\"0\" class=\"header\">Most Used Languages</text>\n        </g>\n    </g>\n\n    <g transform=\"translate(0, 55)\">\n        <svg x=\"25\">\n            <mask id=\"rect-mask\">\n                <rect x=\"0\" y=\"0\" width=\"250\" height=\"8\" fill=\"white\" rx=\"5\"/>\n            </mask>\n            \n                <rect mask=\"url(#rect-mask)\" x=\"0\" y=\"0\" width=\"125\" height=\"8\" fill=\"#00ADD8\"/>\n            \n                <rect mask=\"url(#rect-mask)\" x=\"125\" y=\"0\" width=\"125\" height=\"8\" fill=\"#4F5D95\"/>\n            \n\n            <g transform=\"translate(0, 25)\">\n                <g transform=\"translate(0, 0)\">\n                    \n                        <g transform=\"translate(0, 0)\">\n                            <g class=\"stagger\" style=\"animation-delay: 450ms\">\n                                <circle cx=\"5\" cy=\"6\" r=\"5\" fill=\"#00ADD8\" />\n                                <text x=\"15\" y=\"10\" class='lang-name'>\n                                    Go 50.00%\n                                </text>\n                            </g>\n                        </g>\n                    \n                </g>\n                <g transform=\"translate(150, 0)\">\n                    \n                        <g transform=\"translate(0, 0)\">\n                            <g class=\"stagger\" style=\"animation-delay: 450ms\">\n                                <circle cx=\"5\" cy=\"6\" r=\"5\" fill=\"#4F5D95\" />\n                                <text x=\"15\" y=\"10\" class='lang-name'>\n                                    PHP 50.00%\n                                </text>\n                            </g>\n                        </g>\n                    \n                </g>\n            </g>\n        </svg>\n    </g>\n</svg>\n", res)
	})

	t.Run("Generate main", func(t *testing.T) {
		main()
	})
}
