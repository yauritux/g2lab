package com.yauritux;

import java.math.BigDecimal;
import java.util.Scanner;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

import com.yauritux.model.entity.Card;
import com.yauritux.service.command.CardCommandService;
import com.yauritux.service.query.FareQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Component
@Profile("dev")
public class ApplicationRunner implements CommandLineRunner {
	
	private CardCommandService cardCommandService;
	
	private FareQueryService fareQueryService;
	
	
	
	@Autowired
	void setCardCommandService(CardCommandService cardCommandService) {
		this.cardCommandService = cardCommandService;
	}
	
	@Autowired
	void setFareQueryService(FareQueryService fareQueryService) {
		this.fareQueryService = fareQueryService;
	}
	
	@Override
	public void run(String... arg0) throws Exception {
		
		Scanner scanner = new Scanner(System.in);
		System.out.println("Please enter your name to initialize your new card:");
		String name = scanner.nextLine();
		
		System.out.println("Hi " + name + ". Now please enter the amount as your initial balance:");
		BigDecimal initialBalance = BigDecimal.valueOf(scanner.nextDouble());
		
		Card card = cardCommandService.initializeCard(name, initialBalance);
		
		System.out.println("Congratulations! We have created one card with serial number " + 
				card.getSerialNo() + " for you. Your current balance is £" + card.getBalance());
		
		System.out.println("Would you like to travelling now ? (1) Yes (2) No: ");
		int answer = scanner.nextInt();
		
		boolean travelling = (answer == 1 ? true : false);
		
		while (travelling) {
			System.out.println("What kind of journey you'd like to do ? (1) Bus (2) Tube: ");
			answer = scanner.nextInt();
			switch (answer) {
			case 1:
				//List<Fare> fares = fareQueryService.findByJourneyType(JourneyType.BUS).ge
				break;
			case 2:
				break;
			default:
			}
		}
		
		System.out.println("Thanks for using our service. Have a great day!");
		
		if (scanner != null) {
			scanner.close();
		}
		
		System.exit(0);
	}
}
